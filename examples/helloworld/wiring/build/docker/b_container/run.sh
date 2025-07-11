#!/bin/bash

WORKSPACE_NAME="b_container"
WORKSPACE_DIR=$(pwd)

usage() { 
	echo "Usage: $0 [-h]" 1>&2
	echo "  Required environment variables:"
	
	if [ -z "${B_SERVICE_HTTP_BIND_ADDR+x}" ]; then
		echo "    B_SERVICE_HTTP_BIND_ADDR (missing)"
	else
		echo "    B_SERVICE_HTTP_BIND_ADDR=$B_SERVICE_HTTP_BIND_ADDR"
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


b_process() {
	cd $WORKSPACE_DIR
	
	if [ -z "${B_SERVICE_HTTP_BIND_ADDR+x}" ]; then
		if ! b_service_http_bind_addr; then
			return $?
		fi
	fi

	run_b_process() {
		
        cd b_process
        ./b_process --b_service.http.bind_addr=$B_SERVICE_HTTP_BIND_ADDR &
        B_PROCESS=$!
        return $?

	}

	if run_b_process; then
		if [ -z "${B_PROCESS+x}" ]; then
			echo "${WORKSPACE_NAME} error starting b_process: function b_process did not set B_PROCESS"
			return 1
		else
			echo "${WORKSPACE_NAME} started b_process"
			return 0
		fi
	else
		exitcode=$?
		echo "${WORKSPACE_NAME} aborting b_process due to exitcode ${exitcode} from b_process"
		return $exitcode
	fi
}


run_all() {
	echo "Running b_container"

	# Check that all necessary environment variables are set
	echo "Required environment variables:"
	missing_vars=0
	if [ -z "${B_SERVICE_HTTP_BIND_ADDR+x}" ]; then
		echo "  B_SERVICE_HTTP_BIND_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  B_SERVICE_HTTP_BIND_ADDR=$B_SERVICE_HTTP_BIND_ADDR"
	fi
		

	if [ "$missing_vars" -gt 0 ]; then
		echo "Aborting due to missing environment variables"
		return 1
	fi

	b_process
	
	wait
}

run_all
