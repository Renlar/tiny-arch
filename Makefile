WD=$(shell pwd)
ROOT="img/root.x86_64/"

default: all

all: clean build

build: tar docker

tar: gen_img clean_img set_perm pack_tar rm_img

clean:
	rm -rf root.x86_64.tar.*

gen_img:
	sh $(WD)/arch_gen.sh

clean_img:
	rm -rf $(ROOT)README $(ROOT)etc/resolv.conf $(ROOT)sys $(ROOT)etc/hosts
#	rm -rf $(ROOT)var/cache/pacman/pkg/
#	rm -rf $(ROOT)usr/share/man $(ROOT)usr/share/info $(ROOT)usr/share/doc $(ROOT)usr/share/locale

set_perm:
	sudo chown -R root:root $(ROOT)

pack_tar:
	sudo tar czvf root.x86_64.tar.gz -C $(ROOT) .
	sudo tar cvfJ root.x86_64.tar.xz -C $(ROOT) .

rm_img:
	sudo rm -rf $(ROOT)

docker:
	docker build .
