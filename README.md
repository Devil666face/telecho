### [Do One Thing and Do It Well](https://en.wikipedia.org/wiki/Unix_philosophy#Do_One_Thing_and_Do_It_Well)

```
            .---.
           /     \
           \.@-@./
           /'\_/'\
          //  _  \\
         | \     )|_
        /'\_.>  <_/ \
Telecho \__/'---'\__/
Analog of echo but in telegram chat
Get input data from linux pipe or args
and send in your telegram chat throw telegram bot

Usage:
  telecho [flags]

Examples:
  telecho "This test message"
  telecho Send alert from telecho
  cat file.txt | telecho
  echo "$VARIABLE" | telecho
  BOT_TOKEN="1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" GROUPS_ID="-4444444444,-5555555555" telecho "Your message"
  telecho --config /path/to/your/config.env
  telecho -c /path/to/your/config.{yaml,yml}

Flags:
  -c, --config string   path to config file (default "./.telecho.env ./telecho.{yml,yaml} ~/.telecho.env ~/telecho.{yml,yaml} ~/.config/telecho/.telecho.env ~/.config/telecho/telecho.{yml,yaml}")
  -h, --help            help for telecho
  -v, --version         version for telecho
```

---

#### Configuration file

`.telecho.env` (must end on .env)

```
BOT_TOKEN="1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
GROUPS_ID="-4444444444,-5555555555"
```

`telecho.yaml` (must end on .yaml or .yml)

```yaml
token: "1111111111:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"
groupsId:
  - "-4444444444"
  - "-5555555555"
```
