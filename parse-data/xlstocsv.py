import pandas as pd
from sys import argv

def main():
    filename = argv[1]
    df = pd.read_excel(filename)
    output_file = ".".join(filename.split(".")[:-1]) + ".csv" if len(argv) < 3 else argv[2]
    df.to_csv(output_file, index=False)

if __name__ == "__main__":
    main()

