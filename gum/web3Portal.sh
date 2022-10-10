gum spin --spinner dot --title "Loading Web3 Portal..." -- sleep 2

SSHFILE=~/.ssh/domestic_system
if [[ ! -f "$SSHFILE" ]]; then
  #run seed phrase generator

  #check for .ssh folder
  STORAGEFOLDER=~/.ssh
  [ ! -d "${STORAGEFOLDER}" ] && mkdir -p ${STORAGEFOLDER} && chown 700 ${STORAGEFOLDER}

  KEYCHOICE=$(gum choose "Create a new private key" "Manually enter a seed phrase")

  if [[ "${KEYCHOICE}" ==  "Create a new private key" ]]; then

    gum spin --spinner dot --title "Generating New Seed Phrase & Private Key..." -- sleep 3

    SEEDPHRASE=$(go run ../libs/wallet/generate.go)
    echo "${SEEDPHRASE}" | melt restore "${SSHFILE}" --seed -

    gum style --border normal --margin "1" --padding "1 2" --border-foreground 212 "Please write down this seed phrase, or risk forever losing your keys:"

    gum style --border normal --margin "1" --padding "1 1" --border-foreground 212 "${SEEDPHRASE}"
  else
    #restore key from seed phrase
    SEEDPHRASE=$(gum input --cursor.foreground 212 --prompt.foreground "#0FF" --prompt "* " \
    --placeholder "Enter Your Seedphrase")
    echo "${SEEDPHRASE}" | melt restore "${SSHFILE}" --seed -

    gum style --border normal --margin "1" --padding "1 2" --border-foreground 212 "Key restored from seed phrase"
  fi
  unset SEEDPHRASE
fi
