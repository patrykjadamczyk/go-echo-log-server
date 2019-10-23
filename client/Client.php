<?php

function log($url, $trace=false, $data=[]) {
    $realObj = (is_object($data) or is_array($data)) ? $data : [$data];
    $realUrl = $url.'/logger';
    if ($trace) {
        $realObj['_trace'] = json_encode(debug_backtrace());
    }
    $s = curl_init();
    curl_setopt($s, CURLOPT_URL, $realUrl);
    curl_setopt($s, CURLOPT_HTTPHEADER, []);
    curl_setopt($s, CURLOPT_TIMEOUT, 30);
    curl_setopt($s, CURLOPT_MAXREDIRS, 20);
    curl_setopt($s, CURLOPT_RETURNTRANSFER, true);
    @curl_setopt($s, CURLOPT_FOLLOWLOCATION, true);
    curl_setopt($s, CURLOPT_POST, true);
    curl_setopt($s, CURLOPT_POSTFIELDS, http_build_query($realObj));
    curl_exec($s);
    curl_close($s);
}
