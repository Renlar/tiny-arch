FROM scratch

MAINTAINER renlar <renlar@liddev.com>

ADD https://github.com/redock/tiny-arch/releases/download/0.2.0/root.x86_64.tar.xz /

RUN pacman -Syu --noconfirm && pacman -Scc --noconfirm
#remove unecessary files
RUN rm -rf /usr/share/doc /usr/share/man /usr/share/info /usr/share/locale /var/cache/pacman/pkg/*
