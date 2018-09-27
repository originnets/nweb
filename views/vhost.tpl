server
{
        listen {{.Port}};
        server_name {{.Server_name}};
        index index.php index.html index.htm;

        root  {{.Root}};
        access_log  /var/log/nginx/{{.Logname}}_access.log  main;
        error_log /var/log/nginx/{{.Logname}}_error.log;

        if ($http_user_agent ~* "qihoobot|Baiduspider|Mediapartners-Google|Adsbot-Google|Yahoo! Slurp China|YoudaoBot|Sosospider|Sogou spider|Sogou web spider|MSNBot|ia_archiver|Tomato Bot|bingbot|MJ12bot|YandexBot|AhrefsBot|BLEXBot|UptimeRobot|Googlebot|Googlebot-Mobile|Googlebot-Image|Feedfetcher-Google" ) {
               return 403;
        }
        location / {
                if (!-f $request_filename)     { set $rule_0 1$rule_0; }
                if (!-d $request_filename)     { set $rule_0 2$rule_0; }
                if ($request_filename !~ "-l") { set $rule_0 3$rule_0; }
                if ($rule_0 = "321")           { rewrite ^/(.*)$ /index.php/$1 last; }
         }


        location ~ ^(.+\.php)(.*)$
        {
                fastcgi_pass  127.0.0.1:9000;
                fastcgi_index index.php;
                fastcgi_split_path_info ^(.+\.php)(/.*)$;
                include fastcgi_params;
                fastcgi_param PATH_INFO $fastcgi_path_info;
                fastcgi_param SCRIPT_FILENAME  $document_root$fastcgi_script_name;
        }

        location ~ .*\.(gif|jpg|jpeg|png|bmp|swf|flv|mp3|wma)$
        {
                expires      30d;
        }

        location ~ .*\.(js|css)$
        {
                expires      12h;
        }
}