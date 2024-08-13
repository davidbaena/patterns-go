import json
import subprocess
import sys

import networkx as nx
from matplotlib import pyplot as plt


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

    # changed_files = read_changed_file(absolute_file_path)
    changed_files = ['patterns-go/' + file for file in changed_files]

    json_file_path = './go-viz/emerge-statistics-and-metrics.json'
    metrics_data = read_json_file(json_file_path)

    metrics_files_changed = {file: metrics_data['local-metrics'][file] for file in changed_files if
                             file in metrics_data['local-metrics']}

    print("File Name".ljust(30),
          "methods-in-file".ljust(10),
          "sloc-in-file".ljust(10),
          "louvain-modularity".ljust(10),
          "fan-in".ljust(10),
          "fan-out".ljust(10))
    for file, metrics in metrics_files_changed.items():
        file = file.split('/')[-1]
        print(file.ljust(30),
              str(metrics['number-of-methods-in-file']).ljust(20),
              str(metrics['sloc-in-file']).ljust(20),
              str(metrics['file_result_dependency_graph_louvain-modularity-in-file']).ljust(10),
              str(metrics['fan-in-dependency-graph']).ljust(10),
              str(metrics['fan-out-dependency-graph']).ljust(10))
