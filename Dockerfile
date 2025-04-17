FROM scratch

COPY gosix /gosix

SHELL ["/bin/sh", "-c"]

# bootstrap shell and linking
COPY gosix /bin/ln
COPY gosix /bin/sh
RUN /bin/ln -fs /gosix /bin/ln
RUN /bin/ln -fs /gosix /bin/sh 

# all other utilities
RUN /bin/ln -s /gosix /bin/cat
RUN /bin/ln -s /gosix /bin/false
RUN /bin/ln -s /gosix /bin/ls
RUN /bin/ln -s /gosix /bin/mkdir
RUN /bin/ln -s /gosix /bin/rm
RUN /bin/ln -s /gosix /bin/sleep
RUN /bin/ln -s /gosix /bin/true

CMD ["/bin/sh"]
