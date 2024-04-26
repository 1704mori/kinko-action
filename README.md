# Kinko Action

This is a GitHub|Gitea Action that interacts with the Kinko API.

## Inputs

- `secret`: The name of the secret to use.
- `api_url`: The URL of the Kinko API. Format: `http(s)://(kinko.example.com|192.168.255.254)(:port)@token`.
- `debug`: Enable debug mode to see plain secret output. Options: `true` or `false`. Default is `false`.

## Usage

Here's an example of how to use this action in your workflow:

```yaml
steps:
- name: Add secrets
  uses: https://github.com/1704mori/kinko-action@main
  with:
    secret: my_secret_name
    api_url: ${{ secrets.KINKO_API_URL }}
    debug: false
```

## Credits
- [go-githubactions](https://github.com/sethvargo/go-githubactions)

## License
The scripts and documentation in this project are released under the [GPL 3.0 License](LICENSE).