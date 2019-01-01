import jinja2
import os

def main():

    # Jinja env
    env = jinja2.Environment(loader=jinja2.FileSystemLoader('.'))
    
    problems = [
        {
            'chapter': '2',
            'problem': 'a',
            'title': 'Implement Motif Enumeration',
            'description': 'Given a collection of strings of DNA, find all motifs (kmers of length k and Hamming distance d from all DNA strings).',
            'url': 'http://rosalind.info/problems/ba2a/'
        },
        #{
        #    'chapter': '2',
        #    'problem': 'b',
        #    'title': 'Find a Median String',
        #    'description': 'Given a set of DNA strings, find a k-mer pattern that minimizes the magnitude of the distance from it to the minimum Hamming distance (closest Hamming neighbor) kmer in each DNA string',
        #    'url': 'http://rosalind.info/problems/ba2b/'
        #},
        #{
        #    'chapter': '2',
        #    'problem': 'c',
        #    'title': 'Find a Profile-most Probable k-mer in a String',
        #    'description': 'Given a profile matrix, find the most probable k-mer to generate the given DNA string.',
        #    'url': 'http://rosalind.info/problems/ba2c/'
        #},
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
