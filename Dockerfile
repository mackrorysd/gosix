FROM scratch

COPY gosix /bin/gosix

SHELL ["/bin/sh", "-c"]

# bootstrap shell and linking
COPY gosix /bin/ln
COPY gosix /bin/sh
RUN /bin/ln -fs /bin/gosix /bin/ln
RUN /bin/ln -fs /bin/gosix /bin/sh

# all other utilities
RUN /bin/ln -s /bin/gosix /bin/basename
RUN /bin/ln -s /bin/gosix /bin/cat
RUN /bin/ln -s /bin/gosix /bin/clear
RUN /bin/ln -s /bin/gosix /bin/dirname
RUN /bin/ln -s /bin/gosix /bin/false
RUN /bin/ln -s /bin/gosix /bin/ls
RUN /bin/ln -s /bin/gosix /bin/mkdir
RUN /bin/ln -s /bin/gosix /bin/rm
RUN /bin/ln -s /bin/gosix /bin/sleep
RUN /bin/ln -s /bin/gosix /bin/tee
RUN /bin/ln -s /bin/gosix /bin/true

CMD ["/bin/sh"]
