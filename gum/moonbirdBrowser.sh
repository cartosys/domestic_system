#!/usr/bin/env bash

MOONBIRD=$(curl https://live---metadata-5covpqijaa-uc.a.run.app/images/5268 --output - | jp2a --size=40x20 --colors -)

OUTPUT=$(gum style \
        --foreground 212 --border-foreground 212 --border double \
        --align center  --margin "1 2" --padding "2 4" \
        "${MOONBIRD}")

echo "${OUTPUT}"
