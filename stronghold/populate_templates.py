import jinja2
import os

def main():

    # Jinja env
    env = jinja2.Environment(loader=jinja2.FileSystemLoader('.'))
    
    problems = [
        {
            'id': 'DNA',
            'title': 'Counting DNA Nucleotides',
            'description': 'Given a DNA string, return a count of each base pair as an array, in the order A, C, G, T',
            'url': 'http://rosalind.info/problems/dna/'
        },
    ]
    
    print("Writing problem boilerplate code")
    
    t = 'template.go.j2'
    for problem in problems:
        contents = env.get_template(t).render(**problem)
        fname = problem['id'].lower()+'.go'
        if not os.path.exists(fname):
            print("Writing to file %s..."%(fname))
            with open(fname,'w') as f:
                f.write(contents)
        else:
            print("File %s already exists, skipping..."%(fname))
    
    print("Done")

if __name__=="__main__":
    main()
