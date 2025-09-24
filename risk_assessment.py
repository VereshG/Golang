import os
import openai

# Get OpenAI API key from environment variable
openai.api_key = os.environ.get('OPENAI_API_KEY')

def get_diff_summary():
    # For Jenkins, you can pass the diff summary as an environment variable or read from a file
    # Here, we read from a file called 'pr_diff.txt' (you can change this as needed)
    diff_file = 'pr_diff.txt'
    if os.path.exists(diff_file):
        with open(diff_file, 'r') as f:
            return f.read()
    else:
        return "No diff summary provided."

def assess_risk(diff_summary):
    prompt = f"Assess the risk level of this PR based on the following diff:\n{diff_summary}\nClassify as 'high impact', 'minor change', or 'needs careful review'."
    response = openai.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=[{"role": "user", "content": prompt}]
    )
    return response.choices[0].message.content.strip()

def main():
    diff_summary = get_diff_summary()
    risk = assess_risk(diff_summary)
    print("Risk Assessment:", risk)

if __name__ == "__main__":
    main()
