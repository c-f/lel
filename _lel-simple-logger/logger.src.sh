# Author: Muhammed Taskiran (@0xMillyy); mod by c-f

# Usage 
#export MISATO_OPERATORKEY="lel"
#export MISATO_OPERATORAPI="https://127.0.0.1:8888"
#export MISATO_FILTER="${MISATO_FILTER:-.}"

## --
# FUNCTIONS 

# --[ should logg ?]--
misato_ok(){
    [[ -n "$MISATO_LOGGER" && -n "$MISATO_LOGDIR" ]]
}

# --[toggle logger]--
toggle_misato(){

    case "$1" in 
	on)
		toggle_misato "remote"
		;;
	local)
		export MISATO_LOGGER="L "
		;;
	remote)
		if [ -n "$MISATO_OPERATORAPI" ]; then
		   export MISATO_LOGGER="R "
		else
		   echo "MISATO_OPERATORAPI was not provided !"
		   echo "Fall back to local ..."
		   export MISATO_LOGGER="L "
		fi
		;;
	off)
		export MISATO_LOGGER=""
		;;
	*)
		if [ -n "$MISATO_LOGGER" ]; then
			toggle_misato "off"
		else
			toggle_misato "remote"
		fi
		;;
    esac

}

misato_log(){
	local LEL=$history[$[HISTCMD-1]]
	 if [ -n "$MISATO_OPERATORAPI" ] && [ "$MISATO_LOGGER" = "R " ]; then
		 misato_log_local | filter_command | misato_log_remote
		# misato_log_local | misato_log_remote $LEL
	 else
		 misato_log_local 2>&1 > /dev/null
	 fi
}
milestone_log(){
	if [ -n "$MISATO_OPERATORAPI" ]; then
		milestone_log_local | milestone_log_remote
	else
		milestone_log_local 2> /dev/null
	fi
}



# --[LOCAL LOGGER]--
misato_log_local() {
    get_misato_log | tee -a "${MISATO_LOGDIR}/$(date +%Y-%m-%d)-$(hostname).misato.json"     
}
milestone_log_local() {
    get_milestone_log | tee -a  "${MISATO_LOGDIR}/$(date +%Y-%m-%d)-$(hostname).milestone.json"
}

# --[REMOTE LOGGER]--
milestone_log_remote(){
     send_api_request "/api/core/milestones"
}

misato_log_remote(){
    send_api_request "/api/feat/misato"
}


# --[FUNCTIONS]--
send_api_request(){
	( curl \
            --request POST \
            --header 'Content-Type: application/json' \
            --header "Authorization: Token ${MISATO_OPERATORKEY}" \
            --data-binary @- \
            --no-buffer \
	     -o /dev/null \
	    -s \
	    -k \
	    "${MISATO_OPERATORAPI}$1" &  \
	)
}

filter_command(){
  jq -c \
	  --arg filter "${MISATO_FILTER}" \
	  'select(.command| test("^(" + $filter + ")"))'
}

get_misato_log(){
    jq -nc \
        --arg time "$(date '+%s')" \
        --arg uuid "${MISATO_RUNNING_UUID}" \
        --arg pid "$$" \
        --arg who "$(whoami)" \
        --arg pwd "$PWD" \
        --arg command "$history[$[HISTCMD-1]]" \
        '{time: $time ,termuuid: $uuid, pid: $pid, user: $who, pwd: $pwd, command: $command }'
}
get_milestone_log(){
	 jq -nc \
           --arg milestone "$(zenity --entry --title 'Milestone' --text 'What happend?')" \
           --arg who "$(whoami)" \
           --arg time "$(date '+%s')" \
           '{user: $who, time: $time, milestone: $milestone}'
}


## --
# EXPORTS

export -f misato_ok &> /dev/null
export -f toggle_misato &> /dev/null
export -f misato_log &> /dev/null


## --
# LOGGER

# --[enable Logger]--
# check if misato should be enabled and MISATO is not running
if [[ -n "$MISATO_LOGGER" && -n "$MISATO_LOGDIR" &&  ! -n "$MISATO_RUNNING_UUID" ]]; then
    # test if everything is in place 
    if ! [ -x "$(command -v uuidgen)" ]; then 
        echo "Please install uuidgen"
        return "1"
    fi
    if ! [ -x "$(command -v jq)" ]; then
        echo "Please install jq !"
        return "1" 
    fi
    if ! [ -x "$(command -v curl)" ]; then
	echo "Please install curl !"
	return "1"
    fi

    # save_states
    OLD_HISTTIMEFORMAT="${HISTTIMEFORMAT}"
    HISTTIMEFORMAT=""

    mkdir -p "${MISATO_LOGDIR}"

    # --[set misato]--
    export MISATO_LOGGER=""
    toggle_misato

    export MISATO_RUNNING_UUID="$(uuidgen)"
    export PS1="\${MISATO_LOGGER}[\$(date +%Y-%m-%d_%H:%M:%S)] $PS1"
    export PROMPT_COMMAND=' misato_ok && misato_log;'
    

    if [ -n "$ZSH_VERSION" ]; then
        zsh_source() {
            eval "$PROMPT_COMMAND"
        }
        precmd_functions+=(zsh_source)
    fi
fi
