SOURCE = curl -s http://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv|tr -d '\r'
MODULE = holidays/calendar.go

.PHONY: help
.DEFAULT_GOAL: help

default: clean setup build

###################
# clean
###################

clean: clean_module ## clean

clean_module:
	-rm $(MODULE)

###################
# setup module
###################

setup: clean setup_prefix setup_2016 setup_2017 setup_suffix setup_fmt ## generate holiday calendar

setup_prefix:
	echo "// This file was auto-generated\npackage holidays\nfunc Holidays() []string {\nreturn []string{" >> $(MODULE)

setup_suffix:
	echo "}\n}\n" >> $(MODULE)

setup_2016:
	$(SOURCE)|cut -d, -f2|grep ^2|perl -nle 'print "\"$$_\","' >> $(MODULE)

setup_2017:
	$(SOURCE)|cut -d, -f4|grep ^2|perl -nle 'print "\"$$_\","' >> $(MODULE)

setup_fmt:
	go fmt $(MODULE)

###################
# build
###################

build:
	go build holiday.go
