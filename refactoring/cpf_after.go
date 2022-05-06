package refectoring

import (
	"fmt"
	"strconv"
	"strings"
)

type verificationData struct {
	cpf                                              string
	summationFirstDigit                              int
	summationSecondDigit                             int
	startingNumberSubtractionFirstVerificationDigit  int
	startingNumberSubtractionSecondVerificationDigit int
}

func isInvalidNumberOfCharacters(cpf string) bool {
	var minimumNumberOfCharacters = 11
	var maxmumNumberOfCharacters = 14

	return len(cpf) < minimumNumberOfCharacters || len(cpf) > maxmumNumberOfCharacters
}

func isAllSameDigits(cpf string) bool {
	firstDigit := cpf[0:1]
	var sameDigits bool = true
	for _, currentDigit := range cpf {
		if string(currentDigit) != firstDigit {
			sameDigits = false
			return sameDigits
		}
	}
	return sameDigits
}

func removeSpecialCharacters(cpf string) string {
	var numberCharacters = len(cpf)
	cpf = strings.Replace(cpf, ".", "", numberCharacters)
	cpf = strings.Replace(cpf, "-", "", numberCharacters)
	cpf = strings.Replace(cpf, " ", "", numberCharacters)
	return cpf
}

func firstNineDigits(cpf string) int {
	return len(cpf) - 1
}

func convertCharacterToInteger(character string) (int, error) {
	return strconv.Atoi(character)
}

func computeRestDivision(firstNumber, secondNumber int) int {
	return firstNumber % secondNumber
}

func computeVerificationDigit(restDivision, startingNumberSubtractionVerificationDigit int) int {
	var verificationThreshold int = 2
	if restDivision < verificationThreshold {
		return 0
	}
	return startingNumberSubtractionVerificationDigit - restDivision
}

func computeVerificationDigits(verification verificationData) (int, int, error) {
	var nCount int

	for nCount = 1; nCount < firstNineDigits(verification.cpf); nCount++ {
		currentCharacter := verification.cpf[nCount-1 : nCount]
		currentDigit, err := convertCharacterToInteger(currentCharacter)

		if err != nil {
			return 0, 0, err
		}

		verification.summationFirstDigit += (verification.startingNumberSubtractionFirstVerificationDigit - nCount) * currentDigit
		verification.summationSecondDigit += (verification.startingNumberSubtractionSecondVerificationDigit - nCount) * currentDigit
	}

	return verification.summationFirstDigit, verification.summationSecondDigit, nil
}

func concatenateDigits(firstVerificationDigit, secondVerificationDigit int) string {
	return fmt.Sprintf("%d%d", firstVerificationDigit, secondVerificationDigit)
}

func getLastTwoDigitsFromCPF(cpf string) string {
	return cpf[len(cpf)-2:]
}

func IsValid(cpf string) bool {

	isEmptyCPF := cpf == ""
	if isEmptyCPF {
		return false
	}

	if isInvalidNumberOfCharacters(cpf) {
		return false
	}

	cpf = removeSpecialCharacters(cpf)
	if isAllSameDigits(cpf) {
		return false
	}

	var numberDigitVerification, numberDigitResult string
	var summationFirstDigit, summationSecondDigit, firstVerificationDigit, secondVerificationDigit, restDivision int
	var startingNumberSubtractionFirstVerificationDigit int = 11
	var startingNumberSubtractionSecondVerificationDigit int = 12
	summationFirstDigit = 0
	summationSecondDigit = 0
	firstVerificationDigit = 0
	secondVerificationDigit = 0
	restDivision = 0
	verification := verificationData{
		cpf:                  cpf,
		summationFirstDigit:  summationFirstDigit,
		summationSecondDigit: summationSecondDigit,
		startingNumberSubtractionFirstVerificationDigit:  startingNumberSubtractionFirstVerificationDigit,
		startingNumberSubtractionSecondVerificationDigit: startingNumberSubtractionSecondVerificationDigit,
	}
	summationFirstDigit, summationSecondDigit, err := computeVerificationDigits(verification)
	if err != nil {
		return false
	}
	restDivision = computeRestDivision(summationFirstDigit, startingNumberSubtractionFirstVerificationDigit)
	firstVerificationDigit = computeVerificationDigit(restDivision, startingNumberSubtractionFirstVerificationDigit)
	summationSecondDigit += 2 * firstVerificationDigit
	restDivision = computeRestDivision(summationSecondDigit, startingNumberSubtractionFirstVerificationDigit)
	secondVerificationDigit = computeVerificationDigit(restDivision, startingNumberSubtractionFirstVerificationDigit)
	numberDigitVerification = getLastTwoDigitsFromCPF(cpf)
	numberDigitResult = concatenateDigits(firstVerificationDigit, secondVerificationDigit)
	return numberDigitVerification == numberDigitResult
}
