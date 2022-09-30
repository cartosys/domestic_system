#!/usr/bin/env bash

#set up tmp file system
STORAGEFOLDER=/tmp/domestic_system_storage
[ ! -d "${STORAGEFOLDER}" ] && mkdir ${STORAGEFOLDER}

TOKENID=5268

IMAGEFILE=${STORAGEFOLDER}/${TOKENID}.image
if [[ ! -f "$IMAGEFILE" ]]; then
    curl https://live---metadata-5covpqijaa-uc.a.run.app/images/${TOKENID} > ${IMAGEFILE}
fi
MOONBIRD=$(jp2a --fill --size=40x20 --colors ${IMAGEFILE})

METADATAFILE=${STORAGEFOLDER}/${TOKENID}.data
if [[ ! -f "$METADATAFILE" ]]; then
    curl https://live---metadata-5covpqijaa-uc.a.run.app/metadata/${TOKENID} > ${METADATAFILE}
fi
MOONBIRDDATA=$(cat ${METADATAFILE})

NAME=$(cat ${METADATAFILE} | jq '.name')
FORMATTEDNAME=$(echo "{{ Bold ${NAME} }}" \
    | gum format -t template)

HEADER=$(gum style \
        --foreground 212 --border-foreground 212 --border double \
        --align center  --margin "1 55" --padding "1 1" \
        "Moonbird: ${FORMATTEDNAME}")
OUTPUT=$(gum style \
        --foreground 212 --border-foreground 212 --border double \
        --align center  --margin "1 40" --padding "2 4" \
        "${MOONBIRD}")

echo "${HEADER}"
echo "${OUTPUT}"
echo "${MOONBIRDDATA}"
