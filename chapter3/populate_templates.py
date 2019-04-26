import jinja2
import os

def main():

    # Jinja env
    env = jinja2.Environment(loader=jinja2.FileSystemLoader('.'))
    
    problems = [
        {
            'chapter': '3',
            'problem': 'a',
            'title': 'Generate k-mer Composition of a String',
            'description': 'Given an input string, generate a list of all kmers that are in the input string.',
            'url': 'http://rosalind.info/problems/ba3a/'
        },
        {
            'chapter': '3',
            'problem': 'b',
            'title': 'Reconstruct string from genome path',
            'description': 'Reconstruct a string from its genome path, i.e., sequential fragments of overlapping DNA.',
            'url': 'http://rosalind.info/problems/ba3b/'
        },
        {
            'chapter': '3',
            'problem': 'c',
            'title': 'Construct the overlap graph of a set of k-mers',
            'description': 'Given a set of overlapping k-mers, construct the overlap graph and print a sorted adjacency matrix',
            'url': 'http://rosalind.info/problems/ba3c/'
        },
        {
            'chapter': '3',
            'problem': 'd',
            'title': 'Construct the DeBruijn graph of a string',
            'description': 'Given a DNA input string, construct the DeBruijn graph of all overlapping (k-1)mers',
            'url': 'http://rosalind.info/problems/ba3d/'
        },
        {
            'chapter': '3',
            'problem': 'e',
            'title': 'Construct the DeBruijn graph of a collection of kmers',
            'description': 'Given a collection of arbitrary kmers, construct the DeBruijn graph of all overlapping (k-1)mers',
            'url': 'http://rosalind.info/problems/ba3e/'
        },
    ]
    
    print("Writing problem boilerplate code")
    
    t = 'template.go.j2'
    for problem in problems:
        contents = env.get_template(t).render(**problem)
        fname = 'ba'+problem['chapter']+problem['problem']+'.go'
        if not os.path.exists(fname):
            print("Writing to file %s..."%(fname))
            with open(fname,'w') as f:
                f.write(contents)
        else:
            print("File %s already exists, skipping..."%(fname))
    
    print("Done")

if __name__=="__main__":
    main()
