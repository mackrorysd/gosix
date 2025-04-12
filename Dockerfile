FROM scratch

COPY gosix /gosix

# bootstrap shell and linking
COPY gosix /bin/ln
COPY gosix /bin/sh 
RUN /bin/ln -fs /gosix /bin/ln
RUN /bin/ln -fs /gosix /bin/sh 

CMD ["/bin/sh"]
