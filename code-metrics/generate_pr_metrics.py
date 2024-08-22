import json
import subprocess
import sys
import pandas as pd
import networkx as nx
from matplotlib import pyplot as plt


def json_to_dataframe(json_data):
    # Convert JSON data to a list of dictionaries
    data_list = []
    for file, metrics in json_data.items():
        metrics['file'] = file
        data_list.append(metrics)

    # Convert the list of dictionaries to a pandas DataFrame
    df = pd.DataFrame(data_list)
    return df


def read_json_file(file_path):
    # Read the JSON file into a variable
    with open(file_path, 'r') as file:
        data = json.load(file)
    return data


def get_pr_diff(pr_url, repo):
    # Execute the command and capture the output
    result = subprocess.run(
        ['gh', 'pr', 'diff', pr_url, '--repo', repo, '--name-only'],
        capture_output=True,
        text=True
    )
    # Return the output
    return result.stdout.strip().split('\n')


def plot_nodes_changed(changed_files):
    # Create a directed graph
    # Path to the .graphml file
    graphml_file_path = '../go-viz/emerge-file_result_dependency_graph.graphml'
    # Read the .graphml file
    graph = nx.read_graphml(graphml_file_path)
    # Filter nodes that exist in the lines variable
    filtered_nodes = [node for node in graph.nodes if node in changed_files]
    # Create a subgraph with the filtered nodes
    G = graph.subgraph(filtered_nodes)
    # Plot the subgraph
    plt.figure(figsize=(20, 20))
    pos = nx.spring_layout(G)
    nx.draw(G, pos, with_labels=True, node_size=5000, node_color='skyblue')
    plt.show()


if __name__ == '__main__':
    pr_url = sys.argv[1]
    #    pr_url = 'https://github.com/davidbaena/patterns-go/pull/1'
    repo = 'davidbaena/patterns-go'
    print("PR URL: ", pr_url)

    changed_files = get_pr_diff(pr_url, repo)
    print("Fetch changed files Done!")

    # changed_files = read_changed_file(absolute_file_path)
    changed_files = ['patterns-go/' + file for file in changed_files]

    json_file_path = './go-viz/emerge-statistics-and-metrics.json'
    metrics_data = read_json_file(json_file_path)

    metrics_pr_files_changed = {file: metrics_data['local-metrics'][file] for file in changed_files if
                                file in metrics_data['local-metrics']}
    # convert metrics_pr_files_changed to a pandas DataFrame
    df = json_to_dataframe(metrics_pr_files_changed)

    # load dataset_main_metrics.csv to compare the metrics
    df_main = pd.read_csv('./code-metrics/dataset_main_metrics.csv')

    # Merge the two dataframes and apply the operation subtract to diff the metrics
    merged_df = pd.merge(df, df_main, on='file', how='left',suffixes=('_main', '_pr'))

    # Fill NaN values with 0
    merged_df = merged_df.fillna(0)

    # Calculate the difference between the metrics
    merged_df['methods_diff'] = merged_df['number-of-methods-in-file_main'] - merged_df[
        'number-of-methods-in-file_pr']
    #flipped the order of the subtraction to get the correct diff

    merged_df['sloc_diff'] = merged_df['sloc-in-file_main'] - merged_df['sloc-in-file_pr']
    merged_df['louvain_diff'] = merged_df[
                                                                                    'file_result_dependency_graph_louvain-modularity-in-file_main'] - \
                                                                                merged_df[
                                                                                    'file_result_dependency_graph_louvain-modularity-in-file_pr']
    merged_df['fan_in_diff'] = merged_df['fan-in-dependency-graph_main'] - merged_df[
        'fan-in-dependency-graph_pr']
    merged_df['fan_out_diff'] = merged_df['fan-out-dependency-graph_main'] - merged_df[
        'fan-out-dependency-graph_pr']

    # save the merged dataframe to a csv file
    merged_df[['file',
               'methods_diff',
               'sloc_diff',
               'louvain_diff',
               'fan_in_diff',
               'fan_out_diff']].to_csv('pr_metrics.csv', index=False)

    print(merged_df[['file',
                     'methods_diff',
                     'sloc_diff',
                     'louvain_diff',
                     'fan_in_diff',
                     'fan_out_diff']].to_string(index=False))

