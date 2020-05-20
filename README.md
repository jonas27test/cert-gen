# cert-gen
This application automatically generates certificates. It is heavily based on cert-manager. Please see [https://jonasburster.de/posts/k8s/cert-gen/cert-gen/](this post) for more information.


## Idea
Scan Services for annotations and generate certificates based on these annotations. 

## Ease of use vs granularity
How many parameters should you be able to set via annotations?
