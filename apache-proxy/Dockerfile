FROM httpd:2.4-alpine
COPY ./httpd.conf-from-image /usr/local/apache2/conf/httpd.conf
ADD ./server-root/ /usr/local/apache2/htdocs/
