# TLS and HTTPS

## Mission

Understand and apply TLS and HTTPS in the context of professional Go software engineering.

## Prerequisites

- core-10-18

## Mental Model

TLS is an encrypted tunnel between client and server. The server proves its identity with a certificate issued by a trusted CA. All data sent through the tunnel is encrypted and tamper-proof. HSTS tells browsers to always use the tunnel, refusing plaintext connections.

## Visual Model

```text
input -> concept boundary -> behavior -> observable result
```

## Machine View

TLS 1.3 handshake: ClientHello (supported versions, cipher suites) -> ServerHello (chosen version, cipher, certificate) -> Key Exchange (ECDHE) -> Finished (authenticated encryption). After the handshake, application data flows over the encrypted record layer. Each record is encrypted with a session key derived from the key exchange.

## Run Instructions

```bash
go run ./curriculum/modules/10-auth-security/lessons/19-tls-and-https
go test ./curriculum/modules/10-auth-security/lessons/19-tls-and-https
```

## Code Walkthrough

Read the code from top to bottom. Connect each line to the mental model above. The important question is not only what the line does, but why the program needs that line.

## Try It

1. Run the command exactly as shown.
2. Change one input.
3. Predict the output before running again.
4. Explain the result in your own words.

Common mistakes:

- Using self-signed certificates in production without proper certificate pinning.
- Disabling TLS verification in the client: InsecureSkipVerify: true.
- Using outdated TLS versions (TLS 1.0, 1.1) that are vulnerable to protocol downgrade attacks.
- Not setting HSTS headers, leaving users vulnerable to SSL-stripping attacks.

## In Production

TLS is mandatory for all production web services. Browsers mark HTTP sites as 'not secure'. Let's Encrypt provides free TLS certificates. Go's crypto/tls is used by Kubernetes, Docker, and all major Go web frameworks.

## Thinking Questions

1. What problem does this concept solve?
2. What would break if this concept were removed?
3. How would you test the behavior?

## Next Step

Continue to the next item listed in metadata: `core-10-20`.
