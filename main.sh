#!/bin/bash
BRAVE="Brave Browser"
CHROME="Google Chrome"
SAFARI="Safari"
front_app=$(osascript appscript/frontmost_appname.scpt)

if [ "$front_app" = "$BRAVE"".app" ]; then
  browser=$BRAVE
	url=$(osascript appscript/url_from_brave.scpt)
fi
if [ "$front_app" = "$CHROME"".app" ]; then
  browser=$CHROME
	url=$(osascript appscript/url_from_chrome.scpt)
fi
if [ "$front_app" = "$SAFARI"".app" ]; then
  browser=$SAFARI
	url=$(osascript appscript/url_from_safari.scpt)
fi

echo $html | ./gotosakura "$url" "$browser"