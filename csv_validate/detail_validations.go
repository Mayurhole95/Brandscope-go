package csv_validate

import "strings"

func AdditionalDetail1Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidAdditionalDetail
	}
	return nil
}

func AdditionalDetail2Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func AdditionalDetail3Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func AdditionalDetail4Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func AdditionalDetail5Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func AvailableMonthsValidations(s string) (err error) {
	if s == "" {
		return errAvailableMonthsEmpty
	}
	// fmt.Println(s)
	dates := strings.Split(s, ",")

	if !Equal(dates, dbMonths) {
		return errAvailableMonthsError
	}

	return nil
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func ProductChanges1Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidProductChanges
	}
	return nil
}

func ProductChanges2Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidProductChanges
	}
	return nil
}
func ProductChanges3Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidProductChanges
	}
	return nil
}

func ProductChanges4Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func ProductChanges5Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func ProductSpecification1Validation(s string) (err error) {

	if isAlphaNumeric(s) != nil {
		return errInvalidProductSpecification
	}
	return nil
}

func ProductSpecification2Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidProductSpecification
	}
	return nil
}

func ProductSpecification3Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidProductSpecification
	}
	return nil
}

func ProductSpecification4Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidProductSpecification
	}
	return nil
}

func ProductSpecification5Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidProductSpecification
	}
	return nil
}

func ProductSpecification11Validation(s string) (err error) {

	if isAlphaNumeric(s) != nil {
		return errInvalidProductSpecification
	}
	return nil
}

func ProductSpecification12Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidProductSpecification
	}
	return nil
}

func ProductSpecification13Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidProductSpecification
	}
	return nil
}

func ProductSpecification14Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidProductSpecification
	}
	return nil
}

func ProductSpecification15Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidProductSpecification
	}
	return nil
}

func ProductSpecification6Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func ProductSpecification7Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func ProductSpecification8Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func ProductSpecification9Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}

func ProductSpecification10Validation(s string) (err error) {
	if isAlphaNumeric(s) != nil {
		return errInvalidData
	}
	return nil
}
