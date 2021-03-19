# Chronicler 

SIFT implementation on a time-serie (stock market data).

This is an experiment that did not give the expected results. I leave it here
for reference but it's not useable in any interesting way.

My intention was to find points of interests in stock charts and be able to
compare them in a way that would let me see the effect of a single event (a
news) on multiple stocks. Rather than coming up with my own algorithm to find
points of interests, I tried using the SIFT algorythm, which I would apply to a
time serie rather than an image.

I through this project I realized that image data was significantly different in
nature from the stock market data. In an image, corners and objects make sharp
changes in colors from a pixel to another. In the stock market, data just
moves smoothly up and down over time, and the changes from one point to the next
one are small.

I'm uploading the code here in its unpolished work-in-progress state.
