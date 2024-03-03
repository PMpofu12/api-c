# api-market-data

This api retrieves current and historical crypto data from the coingecko api. It is an abstraction of the CoinGecko api. You can find more documentation for that api here https://www.coingecko.com/api/documentation.

## Still To Do:
1. Pass the parameters for this GET call "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=10" in the request body to the CoinGecko api. (Maybe use a query builder)
2. Add error codes for all endpoints. (Also add in a 429 error for when there were too many requests.)
3. Add unit, integration and acceptance testing to this api. 
4. Add CI/CD to this repo (including stickers for coverage and successful builds).
5. Add a swagger page somewhere.
6. Add in golang linters and always lint code.
7. Write a proper ReadMe.
8. Find a way and a place to deploy this api (hopefully use docker containers and k8s clusters).
