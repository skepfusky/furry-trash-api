import furryQuotes from './furry-quotes.json';
import housepetsQuotes from './housepets-quotes.json';

const headers = {
  "Access-Control-Allow-Origin": "*",
  "Access-Control-Allow-Methods": "GET",
  "Access-Control-Allow-Headers": "Content-Type",
};

export default {
  fetch(req) {
    try {
      const randomFurryQuote = furryQuotes[Math.floor(Math.random() * quotes.length)];
      const randomHousepetsQuote = housepetsQuotes[Math.floor(Math.random() * quotes.length)];

      const url = new URL(req.url);

      switch (url.pathname) {
        case "/housepets":
          return new Response(JSON.stringify({ quote: randomHousepetsQuote }), {
            headers: { ...headers, "content-type": "application/json" },
          });

        case "/housepets/text":
          return new Response(randomHousepetsQuote, {
            headers: { ...headers, "content-type": "text/plain" },
          })

        case "/text":
          return new Response(randomFurryQuote, {
            headers: { ...headers, "content-type": "text/plain" },
          })

        default:
          return new Response(JSON.stringify({ quote: randomFurryQuote }), {
            headers: { ...headers, "content-type": "application/json" },
          })
      }

    } catch (e) {
      return new Response("uh oh the server spooped", {
        status: 500,
        headers: { ...headers },
      });
    }
  }
}