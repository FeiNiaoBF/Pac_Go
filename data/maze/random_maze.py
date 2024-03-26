import random

def generate_maze(width, height):
    symbols = ['#', '.', 'P', 'G', 'X']
    maze = []
    for _ in range(height):
        row = ''.join(random.choices(symbols, weights=[0.3, 0.4, 0.01, 0.05, 0.05], k=width))
        maze.append(row)
    return maze

def print_maze(maze):
    for row in maze:
        print(row)

def main():
    width = input("Please input the width of the maze: ")
    height = input("Please input the height of the maze: ")
    maze = generate_maze(width, height)
    print_maze(maze)

if __name__ == "__main__":
    main()
