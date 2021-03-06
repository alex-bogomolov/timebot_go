# Timebot

**Timebot-go** is a **Slack** bot for holding everyday survey. The bot asks developers how much did they work that day. It also asks on which project they have been working for.

If developer doesn't reply, the bot reminds him about the question. The developer can also log time for any day of the year (except the future days). If a developer didn't log time for previous day, the bot will remind him to report the time in the morning. If a project doesn't exist, a developer can create it. The bot can also provide a time report for a week and since the beginning of the month.

## Requirements
 - Go
 - PostgreSQL

## License

The MIT License (MIT)

Copyright (c) 2017 Codica

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
