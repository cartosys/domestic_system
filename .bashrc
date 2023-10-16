#For cli beautification add the following to ~/.bashrc
#
#sudo apt install lolcat
#sudo apt install figlet

lol()
{
    if [ -t 1 ]; then
        "$@" | lolcat -F .05 -S 290 #toggle last param to gradiate color palette
    else
        "$@"
    fi
}

COMMANDS=(
   ls
   cat
   man
   nano
   git
   ssh
)

for COMMAND in "${COMMANDS[@]}"; do
   alias "${COMMAND}=lol ${COMMAND}"
   alias ".${COMMAND}=$(which ${COMMAND})"
done
echo "$(figlet domestic system)" | lolcat -a -d 1 -F .05 -S 290
#bind 'RETURN: "\e[1~lol \e[4~\n"'  uncomment for all commands
