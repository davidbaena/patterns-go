import json
import pandas as pd


def read_json_file(file_path):
    # Read the JSON file into a variable
    with open(file_path, 'r') as file:
        data = json.load(file)
    return data


def json_to_dataframe(json_data):
    # Convert JSON data to a list of dictionaries
    data_list = []
    for file, metrics in json_data.items():
        metrics['file'] = file
        data_list.append(metrics)

    # Convert the list of dictionaries to a pandas DataFrame
    df = pd.DataFrame(data_list)
    return df


if __name__ == '__main__':
    json_file_path = '../go-viz/emerge-statistics-and-metrics.json'
    metrics_data = read_json_file(json_file_path)
    df = json_to_dataframe(metrics_data['local-metrics'])
    df.to_csv('dataset_main_metrics.csv', index=False)
