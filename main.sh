#!/bin/bash
BRAVE="Brave Browser"
CHROME="Google Chrome"
SIDEKICK="Sidekick"
VIVALDI="Vivaldi"
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
if [ "$front_app" = "$SIDEKICK"".app" ]; then
  browser=$SIDEKICK
	url=$(osascript appscript/url_from_sidekick.scpt)
fi
if [ "$front_app" = "$VIVALDI"".app" ]; then
  browser=$VIVALDI
	url=$(osascript appscript/url_from_vivaldi.scpt)
fi
if [ "$front_app" = "$SAFARI"".app" ]; then
  browser=$SAFARI
	url=$(osascript appscript/url_from_safari.scpt)
fi

echo $html | ./gotosakura "$url" "$browser"