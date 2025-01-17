package parser

import (
	"fmt"
	"github.com/datadog/threatest/pkg/threatest"
	"github.com/datadog/threatest/pkg/threatest/detonators"
	"github.com/datadog/threatest/pkg/threatest/matchers/datadog"
	"sigs.k8s.io/yaml" // we use this library as it provides a handy "YAMLToJSON" function
	"strings"
	"time"
)

// Parse turns a YAML input string into a list of Threatest scenarios
// TODO: A SSH configuration shouldn't be required at this point
func Parse(yamlInput []byte, sshHostname string, sshUsername string, sshKey string) ([]*threatest.Scenario, error) {
	jsonInput, err := yaml.YAMLToJSON(yamlInput)
	if err != nil {
		return nil, fmt.Errorf("unable to convert input YAML to JSON: %v", err)
	}

	parsed := ThreatestSchemaJson{}
	if err := parsed.UnmarshalJSON(jsonInput); err != nil {
		return nil, fmt.Errorf("unable to parse input: %v", err)
	}

	return buildScenarios(&parsed, sshHostname, sshUsername, sshKey)
}

func buildScenarios(parsed *ThreatestSchemaJson, sshHostname string, sshUsername string, sshKey string) ([]*threatest.Scenario, error) {
	scenarios := []*threatest.Scenario{}
	if len(parsed.Scenarios) == 0 {
		return nil, fmt.Errorf("input file has no scenarios defined")
	}

	for _, parsedScenario := range parsed.Scenarios {
		scenario := threatest.Scenario{}
		scenario.Name = parsedScenario.Name

		if !hasDetonation(parsedScenario) {
			return nil, fmt.Errorf("scenario '%s' has no detonation defined", parsedScenario.Name)
		}

		// Detonation
		if localDetonator := parsedScenario.Detonate.LocalDetonator; localDetonator != nil {
			commandToRun := strings.Join(parsedScenario.Detonate.LocalDetonator.Commands, "; ")
			scenario.Detonator = detonators.NewCommandDetonator(&detonators.LocalCommandExecutor{}, commandToRun)
		} else if remoteDetonator := parsedScenario.Detonate.RemoteDetonator; remoteDetonator != nil {
			commandToRun := strings.Join(remoteDetonator.Commands, "; ")
			//TODO: decouple
			//TODO: confirm 1 SSH executor per attack makes sense
			sshExecutor, err := detonators.NewSSHCommandExecutor(sshHostname, sshUsername, sshKey)
			if err != nil {
				return nil, fmt.Errorf("invalid SSH detonator configuration: %v", err)
			}
			scenario.Detonator = detonators.NewCommandDetonator(sshExecutor, commandToRun)
		} else if stratusRedTeamDetonator := parsedScenario.Detonate.StratusRedTeamDetonator; stratusRedTeamDetonator != nil {
			scenario.Detonator = detonators.StratusRedTeamTechnique(*stratusRedTeamDetonator.AttackTechnique)
		}

		// Assertions
		if len(parsedScenario.Expectations) == 0 {
			return nil, fmt.Errorf("scenario '%s' has no assertions defined", parsedScenario.Name)
		}
		for _, parsedAssertion := range parsedScenario.Expectations {
			if datadogMatcher := parsedAssertion.DatadogSecuritySignal; datadogMatcher != nil {
				assertion := datadog.DatadogSecuritySignal(datadogMatcher.Name)
				if severity := datadogMatcher.Severity; severity != nil {
					assertion.WithSeverity(*severity)
				}
				scenario.Assertions = append(scenario.Assertions, assertion)
			}
		}

		//TODO: in the threatest core, the timeout should be part of each assertion (not scenario level)
		// We should probably define a default timeout at the CLI level
		rawTimeout := parsedScenario.Expectations[0].Timeout
		parsedDuration, err := time.ParseDuration(rawTimeout)
		if err != nil {
			return nil, fmt.Errorf("scenario '%s' has an invalid timeout '%s': '%v'", parsedScenario.Name, rawTimeout, err)
		}
		scenario.Timeout = parsedDuration

		scenarios = append(scenarios, &scenario)
	}
	return scenarios, nil
}

// hasDetonation returns true if the scenario has at least 1 detonation defined
func hasDetonation(scenario ThreatestSchemaJsonScenariosElem) bool {
	detonations := scenario.Detonate
	return detonations.LocalDetonator != nil || detonations.RemoteDetonator != nil || detonations.StratusRedTeamDetonator != nil
}
