FROM scratch

COPY gosix /gosix

# bootstrap shell and linking
COPY gosix /bin/ln
COPY gosix /bin/sh 
RUN /bin/ln -fs /gosix /bin/ln
RUN /bin/ln -fs /gosix /bin/sh 

# all other utilities
COPY gosix /bin/cat

CMD ["/bin/sh"]
