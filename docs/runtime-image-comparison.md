# Runtime Image Comparison

## Scratch

Pros

- Smallest possible image
- No shell
- Minimal attack surface

Cons

- Difficult debugging
- No CA certificates by default
- No timezone database

Image Size

~5–7 MB

---

## Distroless

Pros

- Small
- Includes certificates
- Production ready
- Secure
- Recommended by Google

Cons

- No shell

Image Size

~15–25 MB

---

## Debian Slim

Pros

- Easy debugging
- Familiar Linux environment

Cons

- Larger image
- Bigger attack surface

Image Size

~70–120 MB

---

## Decision

This project uses:

Distroless (static-debian12)

Reason:

It provides an excellent balance between security, compatibility, and minimal image size while remaining suitable for production deployments.