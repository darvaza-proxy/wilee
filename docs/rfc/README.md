# ACME RFC Documentation

This directory contains the ACME (Automatic Certificate Management Environment)
RFC specifications and extensions relevant to wilee.

## Core Specifications

- **[RFC 5280](rfc5280.txt)** - Internet X.509 PKI Certificate and CRL Profile.
  Foundation for all certificate operations.
- **[RFC 6962](rfc6962.txt)** - Certificate Transparency. Public logging of
  TLS certificates for auditability.
- **[RFC 8555](rfc8555.txt)** - Automatic Certificate Management Environment
  (ACME). The base protocol specification.
- **[RFC 9794](rfc9794.txt)** - Terminology for Post-Quantum Traditional Hybrid
  Schemes. Definitions for PQ certificate transitions.

## Identifier Extensions

- **[RFC 8738](rfc8738.txt)** - ACME IP Identifier Validation Extension. Enables
  certificate issuance for IP addresses.
- **[RFC 9444](rfc9444.txt)** - ACME for Subdomains. Allows clients to obtain
  certificates for subdomain identifiers.

## Challenge Extensions

- **[RFC 8737](rfc8737.txt)** - ACME TLS Application-Layer Protocol Negotiation
  (ALPN) Challenge Extension. Validates domain control using TLS.

## Certificate Extensions

- **[RFC 8739](rfc8739.txt)** - Support for Short-Term, Automatically Renewed
  (STAR) Certificates in ACME.
- **[RFC 8823](rfc8823.txt)** - Extensions to ACME for End-User S/MIME
  Certificates.
- **[RFC 9115](rfc9115.txt)** - ACME Profile for Generating Delegated
  Certificates. Enables CDNs to obtain certificates on behalf of content
  providers.

## Operational Extensions

- **[RFC 8659](rfc8659.txt)** - DNS Certification Authority Authorization (CAA)
  Resource Record. Base CAA specification.
- **[RFC 8657](rfc8657.txt)** - CAA Record Extensions for Account URI and ACME
  Method Binding.
- **[RFC 9447](rfc9447.txt)** - ACME Challenges Using an Authority Token.
- **[RFC 9773](rfc9773.txt)** - ACME Renewal Information (ARI) Extension.
  Provides server suggestions for certificate renewal timing.

## Drafts

See [../drafts/](../drafts/) for draft specifications not yet published as RFCs.

## Sources

All RFCs downloaded from the [RFC Editor](https://www.rfc-editor.org/).
