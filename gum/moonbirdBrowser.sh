#!/usr/bin/env bash
clear
#set up tmp file system
STORAGEFOLDER=/tmp/domestic_system_storage/moonbirdsBrowser
[ ! -d "${STORAGEFOLDER}" ] && mkdir -p ${STORAGEFOLDER}

TOKENID=$1
#VIEWPORTMIDDLE=$(tput cols)
#VIEWPORTMIDDLE=$(( $VIEWPORTMIDDLE/2 ))

if [[ -z ${TOKENID} ]]; then
  echo "Enter moonbird # (press enter for random)"
  TOKENID=$(gum input --cursor.foreground "#FF0" --prompt.foreground "#0FF" --prompt "* " --placeholder 1234)
fi
if [[ -z ${TOKENID} ]]; then
    TOKENID=$(( ( RANDOM % 10000 ) ))
fi

IMAGEFILE=${STORAGEFOLDER}/${TOKENID}.image
if [[ ! -f "$IMAGEFILE" ]]; then
    until curl -s -f -o  ${IMAGEFILE} https://live---metadata-5covpqijaa-uc.a.run.app/images/${TOKENID}
    do
      gum spin --spinner dot --title "Fetching Moonbird image..." -- sleep 4
    done
fi
#MOONBIRD=$(jp2a --fill --size=40x20 --colors ${IMAGEFILE})
MOONBIRD=$(cat ${IMAGEFILE} | imgcat )

METADATAFILE=${STORAGEFOLDER}/${TOKENID}.data
if [[ ! -f "$METADATAFILE" ]]; then
    until curl -s -f -o ${METADATAFILE} https://live---metadata-5covpqijaa-uc.a.run.app/metadata/${TOKENID}
    do
      gum spin --spinner dot --title "Fetching Moonbird data..." -- sleep 4
    done
fi
MOONBIRDDATA=$(cat ${METADATAFILE} | jq -c '.attributes')

NAME=$(cat ${METADATAFILE} | jq '.name')
FORMATTEDNAME=$(echo "{{ Bold ${NAME} }}" \
    | gum format -t template)

HEADER=$(gum style \
        --foreground 212 --border-foreground 212 --border double \
        --align center   --margin "1 55" --padding "1 1" \
        "Moonbird: ${FORMATTEDNAME}")
OUTPUT=$(gum style \
        --foreground 212 --border-foreground 212 --border double \
        --align center  --margin "1 40" --padding "2 4" \
        "${MOONBIRD}")

clear
echo "${HEADER}"
echo "${OUTPUT}"
TRAITTABLE=""

function getTraitTable () {
  for row in $(echo "${MOONBIRDDATA}" | jq -r '.[] | @base64'); do

        KEY=$(echo ${row} | base64 --decode | jq -r '.trait_type')
        SEPERATOR=": "
        VALUE=$(echo ${row} | base64 --decode | jq -r '.value')
        #TRAITTABLE+=$(printf "${KEY}${SEPERATOR}${VALUE}\n\n")
        TRAITTABLE+=$(echo "${KEY}: ${VALUE}")
        echo $(echo "${KEY}: ${VALUE}")
  done
}


FOOTER=$(gum style \
        --foreground 212 --border-foreground 212 --border double \
        --align center --margin "1 50" --padding "2 4" \
        "$(getTraitTable)")
echo "${FOOTER}"
