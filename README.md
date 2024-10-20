# OTP Package

This Go package provides implementations for One-Time Password (OTP) systems, including TOTP (Time-Based One-Time Password), HOTP (HMAC-Based One-Time Password), and OCRA (OATH Challenge-Response Algorithm). It offers a secure, flexible, and easy-to-use solution for enhancing user authentication.

## Features

- **TOTP**: Generates time-based one-time passwords.
- **HOTP**: Generates one-time passwords based on a counter value.
- **OCRA**: Implements the OATH Challenge-Response algorithm for user verification.
- **Ease of Use**: Simple API for quick and easy integration.

## Installation

```bash
go get github.com/hasan-kilici/otp
```

## Usage
### 1. OTP Configuration
```go
package main

import (
	"fmt"
	"log"
	"github.com/hasan-kilici/otp"
	"github.com/hasan-kilici/otp/utils"
)

func main() {
	otpConfig := otp.Config{
		Secret:      "JBSWY3DPEHPK3PXP",
		TimeInterval: 30,
		Digits:      6,
		Algorithm:   utils.AlgorithmSHA1,
	}

	otpOptions := otp.Options{
		Issuer:      utils.DefaultIssuer,
		AccountName: utils.DefaultAccount,
	}

	otpInstance := otp.New(otpConfig)

	otpCode, err := otpInstance.Generate(otpOptions)
	if err != nil {
		log.Fatalf("OTP Error: %v", err)
	}
	fmt.Println("OTP Code:", otpCode)
}
```

Description
  - Secret: The shared secret used for generating OTP.
  - TimeInterval: Duration for which the OTP is valid (in seconds).
  - Digits: Number of digits in the OTP.
  - Algorithm: Hashing algorithm used (e.g., SHA1).

### 2. TOTP Configuration

```go
package main

import (
	"fmt"
	"log"
	"github.com/hasan-kilici/otp/totp"
	"github.com/hasan-kilici/otp/utils"
)

func main() {
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
}
```
Description
- Secret: Shared secret for TOTP generation.
- TimeStep: Time interval (in seconds) for which TOTP is valid.
- Digits: Number of digits in the generated TOTP.
- Algorithm: Hashing algorithm used (e.g., SHA1).

### 3. HOTP Configuration
```go
package main

import (
	"fmt"
	"log"
	"github.com/hasan-kilici/otp/hotp"
	"github.com/hasan-kilici/otp/utils"
)

func main() {
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
}
```
Description

- Secret: Shared secret for generating HOTP.
- Counter: Counter value for HOTP generation.
- Digits: Number of digits in the generated HOTP.
- Algorithm: Hashing algorithm used (e.g., SHA256).

### 4. OCRA Configuration
package main
```go
import (
	"fmt"
	"log"
	"github.com/hasan-kilici/otp/ocra"
	"github.com/hasan-kilici/otp/utils"
)

func main() {
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
```
Description
- SecretKey: Shared secret key for OCRA.
- Suite: The OCRA suite being used (e.g., OCRA-1).
- Counter: Counter value for OCRA generation.
- Algorithm: Hashing algorithm used (e.g., SHA512).

---

### References
- [RFC 4226: HOTP: An HMAC-Based One-Time Password Algorithm](https://www.ietf.org/rfc/rfc4226.txt)
- [RFC 6238: TOTP: Time-Based One-Time Password Algorithm](https://datatracker.ietf.org/doc/html/rfc6238)
- [RFC 6287: OCRA: OATH Challenge-Response Algorithm](https://www.rfc-editor.org/rfc/rfc6287.html)
