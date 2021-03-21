
# Daft Wullie

Daft Wullie is a lightweight markup language for taking notes quickly. Based on Markdown, I intend it for note taking in meetings, lectures, and during research sessions where readability and speed are the most desirable traits.

## Why?

During meetings and lextures I prefer to write notes on paper as it's quicker and less painful than typing on some device. However, when I'm researching by myself or as a team I prefer to type as it allows me to save and share easily. I'm also aware that many other people do like to type rather than write notes in lectures and meetings.

I gave up on the following:
- Plain text is nice and simple but lacks the expressiveness that helps readability later
- Office document formats require office tools, e.g. MS Word, and provide so many features that choosing decoration for future readability distracts me

I've found Markdown to be the best format so far as there are tools that can continuosly update a representation of your notes as you write them, e.g. https://dillinger.io/. Markdown has some drawbacks in this context though:
1. The language is still a little too large for writing notes very quickly, especially in lectures and meetings
2. For simple note taking tasks it overloads my tiny brain with too many syntax rules and choices
3. It doesn't allow me to highlight key words and phrases or document positive and negative points that I would like to see highlighted when I review my notes

If I can tag key phrases, along with positive and negative points, I can create simple tools to compile and analyse them. E.g. a tool that compiles all keywords can be used for the next search target during research. In Markdown I use strong text decoration but this is easily confused with words in a phrase I'm only trying to emphasise, like '**Not**'.

So I decided I need a really simple and concise markup language for quickly writing and sharing readable notes.

## Example
```
# Cheese
"Cheese is a dairy product, derived from milk and produced in wide ranges of flavors, textures and forms by coagulation of the milk protein casein." $Wikipedia
Cheese is +very tasty+ but also quite -smelly-, +good on pizza though+

## History
Who knows.

## Types
. Chedder, always from $Chedder,Somerset,England
. Brie
. Mozzarella
. Stilton
. and many more

## Process
! Curdling
!! Souring
!! Adding Rennet
! Curd processing
.. Stretching
.. Cheddering
.. Washing
! Ripening

## Bacteria
Milk used should be **pasteurized** to kill infectious diseases

## Heart disease
-Recommended that cheese consumption be minimised
-There isn't any *convincing* evidence that cheese lowers heart disease

Source: https://en.wikipedia.org/wiki/Cheese $2021-02-06
```

## Language Rules

1. Each line is parsed independently, **there are no multiline features!** This simplicity is great for both writers and tool implementors. Representation tools may group lines together as they see fit, e.g. to present lists nicely
2. Any symbol placed anywhere can be interpretted in one and only one valid way, i.e. it is impossible to have a syntax error during interpretation so any errors are errors introduced by the user 
3. A backslash `\` is used to escape a symbol

### Line Nodes
- Annotates an entire line
- Allows **phrase nodes** within the line 
- Start with a specific symbol as the first none whitespace character in a line and end at the next linefeed
- A line that does not start with any line node symbol is a text line, which may or may not contain **phrase nodes**

### Phrase Nodes

- Annotates their text content
- Start with a specific symbol and end either at the first linefeed or when its corresponding end symbol is encountered
- Can be nested only if a parent and child pair are of different types (Snippets are the exception, they cannot contain nested nodes)

### Node Table

| Lead symbol | Node Type | Description |
| :---: | :---: | :--- |
| `#` | Line | Topic |
| `##` | Line | Subtopic |
| `.` | Line | An unordered list item |
| `..` | Line | An unordered sub-list item (indented) |
| `!` | Line | An ordered list item |
| `!!` | Line | An ordered sub-list item (indented) |
| `**` | Phrase | Keyword or key phrase |
| `+` | Phrase | Positive phrase |
| `-` | Phrase | Negative phrase |
| `*` | Phrase | Emphasis, bold, or strong phrase |
| ``` ` ``` | Phrase | A snippet of code or literal text, nested nodes are **not** supported |
| `"` | Phrase | A quote |
| `$` | Phrase | Artifact: a person, group, place, or datetime |

### Lists

Ordered and unordered list items cover one line and it is up to the representation tool to bring these lines together into a styled block.

**Example:**

```
## Act like a PID Controller
! Identify target position (might have changed)
! Identify current position (should have changed)
! Calculate trajectory to target
! Take a small step towards the target
! Repeat
```
