.PHONY: $(PB_DIR_TGTS)
$(PB_DIR_TGTS):
	@for dir in $(subst _PB,, $@); do \
		echo Now Build proto in directory: $$dir; \
		cd $$dir; rm -rf mock; \
		export PATH=$(PATH); \
		rm -f *.pb.go; rm -f *.trpc.go; \
		find . -name '*.proto' | xargs -I DD \  # 将它们传递给后续的命令 DD是一个占位符，查找到的文件会替换DD
			trpc create -f --protofile=DD --protocol=trpc --rpconly --nogomod --alias --mock=false --protodir=$(WORK_DIR)/proto; \
		ls *.trpc.go | xargs -I DD mockgen -source=DD -destination=mock/DD -package=mock ; \
		find `pwd` -name '*.pb.go'; \
	done



# 启动服务，指定conf文件
	#cd app/user/; go run . -conf conf/trpc_go.yaml
	#cd app/http-auth-server/; go run . -conf conf/trpc_go.yaml