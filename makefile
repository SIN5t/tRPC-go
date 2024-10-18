.PHONY: $(PB_DIR_TGTS)
$(PB_DIR_TGTS):
	@for dir in $(subst _PB,, $@); do \
		echo Now Build proto in directory: $$dir; \
		cd $$dir; rm -rf mock; \
		export PATH=$(PATH); \
		rm -f *.pb.go; rm -f *.trpc.go; \
		find . -name '*.proto' | xargs -I DD \
			trpc create -f --protofile=DD --protocol=trpc --rpconly --nogomod --alias --mock=false --protodir=$(WORK_DIR)/proto; \
		ls *.trpc.go | xargs -I DD mockgen -source=DD -destination=mock/DD -package=mock ; \
		find `pwd` -name '*.pb.go'; \
	done