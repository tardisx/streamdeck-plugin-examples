# Example Stream Deck plugin - "spinclock"

![Demo](anim.gif)

This plugin displays a minimalist clock - the number on the clock is the hour
(24h time) and the rotation indicates the minute. With a little practice it
should become easy to tell the time with some accuracy.

Tapping on a clock changes its colour to a random colour.

# Trying it out

Check this code out somewhere.

Symlink `the au.id.hawkins.sd.spinclock.sdPlugin` directory into your
plugin directory. See [Elgato's documentation](https://docs.elgato.com/sdk/plugins/getting-started#id-4.-add-the-plugin-to-stream-deck).

Compile the code, using the `./build.sh` script (sorry Windows users, no `.bat` file, patches welcome).

Restart the Stream Deck software, your plugin should now be available in the list on the right hand side. When you drag it
onto your profile, the plugin will start.

Stdout/stderr logs are available, on Mac they are at:

    /Users/<username>/Library/Logs/ElgatoStreamDeck

# Making changes

After modifying the code, rebuild using the script, and simply kill the running process to make Stream Deck restart it for you:

    killall spinclock

Note that if your plugin restarts too many times in short succession,
Stream Deck will disable it completely (see the logs above) - the
only way I know of to recover is to restart the Stream Deck software.
