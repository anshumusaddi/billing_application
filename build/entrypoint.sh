#! /bin/bash

export ENV=${ENV:-local}
export ENV=/app/config/$ENV.yaml

./billing_application