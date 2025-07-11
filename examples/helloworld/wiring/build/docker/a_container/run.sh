#!/bin/bash

WORKSPACE_NAME="a_container"
WORKSPACE_DIR=$(pwd)

usage() { 
	echo "Usage: $0 [-h]" 1>&2
	echo "  Required environment variables:"
	
	if [ -z "${A_SERVICE_HTTP_BIND_ADDR+x}" ]; then
		echo "    A_SERVICE_HTTP_BIND_ADDR (missing)"
	else
		echo "    A_SERVICE_HTTP_BIND_ADDR=$A_SERVICE_HTTP_BIND_ADDR"
	fi
	if [ -z "${B_SERVICE_HTTP_DIAL_ADDR+x}" ]; then
		echo "    B_SERVICE_HTTP_DIAL_ADDR (missing)"
	else
		echo "    B_SERVICE_HTTP_DIAL_ADDR=$B_SERVICE_HTTP_DIAL_ADDR"
	fi
		
	exit 1; 
}

while getopts "h" flag; do
	case $flag in
		*)
		usage
		;;
	esac
done


a_process() {
	cd $WORKSPACE_DIR
	
	if [ -z "${B_SERVICE_HTTP_DIAL_ADDR+x}" ]; then
		if ! b_service_http_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${A_SERVICE_HTTP_BIND_ADDR+x}" ]; then
		if ! a_service_http_bind_addr; then
			return $?
		fi
	fi

	run_a_process() {
		
        cd a_process
        ./a_process --b_service.http.dial_addr=$B_SERVICE_HTTP_DIAL_ADDR --a_service.http.bind_addr=$A_SERVICE_HTTP_BIND_ADDR &
        A_PROCESS=$!
        return $?

	}

	if run_a_process; then
		if [ -z "${A_PROCESS+x}" ]; then
			echo "${WORKSPACE_NAME} error starting a_process: function a_process did not set A_PROCESS"
			return 1
		else
			echo "${WORKSPACE_NAME} started a_process"
			return 0
		fi
	else
		exitcode=$?
		echo "${WORKSPACE_NAME} aborting a_process due to exitcode ${exitcode} from a_process"
		return $exitcode
	fi
}


run_all() {
	echo "Running a_container"

	# Check that all necessary environment variables are set
	echo "Required environment variables:"
	missing_vars=0
	if [ -z "${A_SERVICE_HTTP_BIND_ADDR+x}" ]; then
		echo "  A_SERVICE_HTTP_BIND_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  A_SERVICE_HTTP_BIND_ADDR=$A_SERVICE_HTTP_BIND_ADDR"
	fi
	
	if [ -z "${B_SERVICE_HTTP_DIAL_ADDR+x}" ]; then
		echo "  B_SERVICE_HTTP_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  B_SERVICE_HTTP_DIAL_ADDR=$B_SERVICE_HTTP_DIAL_ADDR"
	fi
		

	if [ "$missing_vars" -gt 0 ]; then
		echo "Aborting due to missing environment variables"
		return 1
	fi

	a_process
	
	wait
}

run_all
