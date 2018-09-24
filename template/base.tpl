# For more information on configuration, see:
#   * Official English Documentation: http://nginx.org/en/docs/
#   * Official Russian Documentation: http://nginx.org/ru/docs/

user www;
worker_processes auto;
# error_log /var/log/nginx/error.log;
pid /var/run/nginx.pid;

events {
    worker_connections  1024;
}


http {

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

#    access_log  /var/log/nginx/access.log  main;

    sendfile            on;
    tcp_nopush          on;
    tcp_nodelay         on;
    gzip                on;
    keepalive_timeout   60 60;
    fastcgi_connect_timeout 300;
    fastcgi_send_timeout 300;
    fastcgi_read_timeout 300;
    fastcgi_buffer_size 16k;
    fastcgi_buffers 16 16k;
    fastcgi_busy_buffers_size 16k;
    fastcgi_temp_file_write_size 16k;
    fastcgi_intercept_errors on;



    client_header_buffer_size 4k;
    large_client_header_buffers 4 4k;
    client_max_body_size 50m;
    types_hash_max_size 2048;
    charset utf-8;
    server_tokens off;

    include             mime.types;
    default_type        application/octet-stream;

    # Load modular configuration files from the /etc/nginx/conf.d directory.
    # See http://nginx.org/en/docs/ngx_core_module.html#include
    # for more information.
    include /usr/local/nginx/vhost/*.conf;

}