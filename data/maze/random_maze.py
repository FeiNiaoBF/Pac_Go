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
    for _ in range(5):
        width, height = 25, 10  # 设定迷宫的宽度和高度
        maze = generate_maze(width, height)
        print_maze(maze)
        print()  # 每个迷宫之间空一行

if __name__ == "__main__":
    main()
