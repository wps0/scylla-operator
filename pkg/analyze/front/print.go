package front

import (
	"errors"
	"fmt"

	"github.com/scylladb/scylla-operator/pkg/analyze/symptoms"
)

func Print(issue symptoms.Issue, compact bool) error {
	if issue.Symptom == nil {
		return errors.New("invalid argument: symptom cannot be nil")
	}

	fmt.Println((*issue.Symptom).Name())

	if compact {
		return nil
	}

	diagnoses := (*issue.Symptom).Diagnoses()
	if len(diagnoses) == 0 {
		fmt.Println("No diagnoses found for this issue.")
	} else {
		fmt.Println("Diagnoses:")
		for _, diagnosis := range diagnoses {
			fmt.Println("\t", diagnosis)
		}
	}

	suggestions := (*issue.Symptom).Suggestions()
	if len(suggestions) == 0 {
		fmt.Println("No suggestions found for this issue.")
	} else {
		fmt.Println("Suggestions:")
		for _, suggestion := range suggestions {
			fmt.Println("\t", suggestion)
		}
	}

	resources := issue.Resources
	if len(resources) == 0 {
		fmt.Println("No resources related to this issue.")
	} else {
		for key := range issue.Resources {
			fmt.Println(key)
		}
	}

	return nil
}
