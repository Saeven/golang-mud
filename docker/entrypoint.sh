#!/bin/sh

ls .
ls -l $APP_HOME
$APP_HOME/$DIST_NAME "$@"
