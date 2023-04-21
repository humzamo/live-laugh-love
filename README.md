# Live, Laugh, Love ✨

An API to document your mood with three easy categories: live, laugh, and love 🌱😂❤️

## What does it mean to live, laugh, love?

- **LIVE** 🌱 - your “zest for life”.
  How motivated are you feeling? Are you feeling jovial? How excited are you for the day?

- **LAUGH** 😂 - your humour levels for the day.
  How jokey are you feeling? Is it easy to make you laugh? Are you in the mood to make others laugh?

- **LOVE** ❤️ - your gratitude for the day.
  Are you feeling the positive energy from your peers/colleagues/friends/family/partner(s)? Are you radiating good vibes? Do you feel gratuitous/appreciated?

Each score is recorded on a scale from 0-10 (inclusive). We recommend entering a score at least once per day and taking a few minutes to reflect on how they impact your mood.

## How to run the api

To run the API, simply run the following command in your terminal to run the GQL server:

```
make run-api
```

You can make query (such as via PostMan) using the following syntax:

```
mutation examplePost($message: String = "hello!") {
  examplePost(message: $message)
}
```

```
query exampleGet {
  exampleGet
}
```

```
query recordByUser($user: String = "your name", $date: Date = "2023-01-01") {
  recordByUser(user: $user, date: $date) {
    user,
    date,
    reason,
    scores{
            live,
            laugh,
            love,
    },
  }
}
```
