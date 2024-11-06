<script>
  export let name;

  import { GreeterPromiseClient } from "./helloworld/proto/helloworld_grpc_web_pb";
  import { HelloRequest } from "./helloworld/proto/helloworld_pb";

  const client = new GreeterPromiseClient("http://localhost:51051", null, null);
  const request = new HelloRequest();
  request.setName("Tanuki");
  client.sayHello(request, {}).then((reply) => {
    console.log(reply.getMessage());
    name = reply.getMessage();
  });
</script>

<main>
  <h1>{name}!</h1>
  <p>
    Visit the <a href="https://svelte.dev/tutorial">Svelte tutorial</a> to learn
    how to build Svelte apps.
  </p>
</main>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
