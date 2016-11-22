FROM ubuntu

RUN export GOPATH=$HOME/go && \
    echo export PATH=$PATH:$GOPATH/bin && \
    bash && \
    apt-get update && \
    apt-get install -y git && \
    apt-get install -y golang && \
    apt-get install -y bash-completion && \
    go get -u github.com/ebsarr/packet && \
    mv $GOPATH/bin/packet /usr/local/bin && \
    /usr/local/bin/packet genautocomplete && \
    mv packet-autocomplete.sh /etc/bash_completion.d/packet &&\
    sed -i -e 's/\#if \[ -f \/etc\/bash_completion/if \[ -f \/etc\/bash_completion/' -e 's/\#    \. \/etc\/bash_completion/    \. \/etc\/bash_completion/' -e 's/\#fi$/fi/' $HOME/.bashrc &&\
    rm -rf $HOME/go &&\
    apt-get remove -y --purge golang &&\
    apt-get remove -y --purge git

CMD /usr/local/bin/packet -h