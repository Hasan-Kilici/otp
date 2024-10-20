/*
Package otp provides an implementation of one-time password (OTP) generation using TOTP, HOTP, and OCRA algorithms.

The OTP generation is based on the following RFCs:

- RFC 4226: HOTP: An HMAC-Based One-Time Password Algorithm
- RFC 6238: TOTP: Time-Based One-Time Password Algorithm
- RFC 6287: OCRA: OATH Challenge-Response Algorithm

## Overview

This package contains structures and methods for generating OTPs using different algorithms:

1. **TOTP** (Time-based One-Time Password)
2. **HOTP** (HMAC-based One-Time Password)
3. **OCRA** (OATH Challenge-Response Algorithm)

## Usage

To use the package, create a configuration for the desired OTP type, set the necessary parameters, and call the Generate method to obtain the OTP.

### Example

package main

import (
	"fmt"
	"log"
	"github.com/hasan-kilici/otp/totp"
	"github.com/hasan-kilici/otp/ocra"
	"github.com/hasan-kilici/otp/hotp"
	"github.com/hasan-kilici/otp/utils"
)

func main() {
	// TOTP Example
	totpConfig := totp.Config{
		Secret:    "JBSWY3DPEHPK3PXP",
		TimeStep:  30,
		Digits:    6,
		Algorithm: utils.AlgorithmSHA1,
	}
	totpOptions := totp.Options{
		Issuer:      utils.DefaultIssuer,
		AccountName: utils.DefaultAccount,
	}
	totpInstance := totp.New(totpConfig)
	totpCode, err := totpInstance.Generate(totpOptions)
	if err != nil {
		log.Fatalf("TOTP Error: %v", err)
	}
	fmt.Println("TOTP Code:", totpCode)

	// HOTP Example
	hotpConfig := hotp.Config{
		Secret:    "JBSWY3DPEHPK3PXP",
		Counter:   1,
		Digits:    6,
		Algorithm: utils.AlgorithmSHA256,
	}
	hotpOptions := hotp.Options{
		Issuer:      utils.DefaultIssuer,
		AccountName: utils.DefaultAccount,
	}
	hotpInstance := hotp.New(hotpConfig)
	hotpCode, err := hotpInstance.Generate(hotpOptions)
	if err != nil {
		log.Fatalf("HOTP Error: %v", err)
	}
	fmt.Println("HOTP Code:", hotpCode)

	// OCRA Example
	ocraConfig := ocra.Config{
		SecretKey:   "JBSWY3DPEHPK3PXP",
		Suite:       "OCRA-1",
		Counter:     "1",
		Algorithm:   utils.AlgorithmSHA512,
	}
	ocraOptions := ocra.Options{
		Issuer:      utils.DefaultIssuer,
		AccountName: utils.DefaultAccount,
	}
	ocraInstance := ocra.New(ocraConfig)
	ocraCode, err := ocraInstance.Generate(ocraOptions)
	if err != nil {
		log.Fatalf("OCRA Error: %v", err)
	}
	fmt.Println("OCRA Code:", ocraCode)
}

## Structures

### Config

The Config structure contains the settings necessary for OTP generation:

- Secret: The shared secret used for generating OTPs.
- TimeStep: The time interval for TOTP (in seconds).
- Counter: The counter value for HOTP.
- Digits: The number of digits in the generated OTP.
- Algorithm: The hash algorithm used (SHA1, SHA256, SHA512).

### Options

The Options structure contains optional parameters for OTP generation:

- Issuer: The name of the application issuing the OTP.
- AccountName: The account name for which the OTP is generated.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.
*/
package otp
