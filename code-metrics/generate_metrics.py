import json
import networkx as nx
from matplotlib import pyplot as plt


def read_json_file(file_path):
    # Read the JSON file into a variable
    with open(file_path, 'r') as file:
        data = json.load(file)
    return data


def read_changed_file(file_path):
    # Read the file into a list
    with open(file_path, 'r') as file:
        lines = [line.strip() for line in file]
    return lines


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
    # Define the file path
    absolute_file_path = '../changed_files.txt'

    changed_files = read_changed_file(absolute_file_path)
    changed_files = ['patterns-go/' + file for file in changed_files]

    json_file_path = '../go-viz/emerge-statistics-and-metrics.json'
    metrics_data = read_json_file(json_file_path)

    metrics_files_changed = {file: metrics_data['local-metrics'][file] for file in changed_files if
                             file in metrics_data['local-metrics']}

    print("File Name".ljust(50),
          "methods-in-file".ljust(10),
          "sloc-in-file".ljust(10),
          "louvain-modularity".ljust(10),
          "fan-in".ljust(10),
          "fan-out".ljust(10))
    for file, metrics in metrics_files_changed.items():
        print(file.ljust(50),
              str(metrics['number-of-methods-in-file']).ljust(20),
              str(metrics['sloc-in-file']).ljust(20),
              str(metrics['file_result_dependency_graph_louvain-modularity-in-file']).ljust(10),
              str(metrics['fan-in-dependency-graph']).ljust(10),
              str(metrics['fan-out-dependency-graph']).ljust(10))


# %%
