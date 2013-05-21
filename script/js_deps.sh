#!/bin/bash

set -e

mkdir -p public/js/lib

wget -O public/js/lib/jquery.js http://code.jquery.com/jquery-2.0.0.js
wget -O public/js/lib/handlebars.js https://raw.github.com/wycats/handlebars.js/1.0.0-rc.4/dist/handlebars.js
wget -O public/js/lib/ember.js http://builds.emberjs.com.s3.amazonaws.com/ember-latest.js
wget -O public/js/lib/ember-data.js http://builds.emberjs.com.s3.amazonaws.com/ember-data-latest.js
