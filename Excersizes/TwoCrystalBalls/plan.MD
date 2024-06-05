# Two Crystal Balls

Problem: `Given two crystal balls that will break if dropped from a high enough distance, determine the exact spot 
which it will break in the most optimal way.`

Plan:
We have two crystal balls which means we have two chances at finding the point at which the crystal balls breaks.
Initially we realize we could just simply go from one floor to the next dropping the crystal ball until it breaks,
although that's correct it is not effecient as what if we have to run all the way to the top floor to finally break
it, that's a whole lota stairs.

What about binary search? Well we could take the elevator to the middle of the building and drop it out of the window.
But shoot, it just broke and we only have one left, can we risk going down half as many floors and trying again? No.
So that means now we have to resort back to our original idea.

Whats one way we can more efficiently test our idea? Well instead of walking N each time what if we jumped forward sqrt(N)
Now when the crystal ball breaks we only have to go back one sqrt(N) instead of all the way back to 0 and we can
walk forward linearly meainging our time complexity drops to O(2 sqrt(N))

So now lets implement.

First we will create a method MinBreakingPoint that takes in our array of booleans that when true will bust our
balls or false will leave them unscaved.

First we will iterate through the array(slice). Our first step is to check if the crystal ball has broken at the
current index. If false we jump forward one sqrt(N) and we check again. If true we go back to the last know safest height. If the last known safest height is equal to the
current index we return the index. If not, we then jump back and linearly search this subset between the last known
safe point and sqrt(n).


